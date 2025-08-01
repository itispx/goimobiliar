package condom_condominio_consultar

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/itispx/goimobiliar/condominio/bloco"
	"github.com/itispx/goimobiliar/condominio/categoria"
	"github.com/itispx/goimobiliar/condominio/classificacao"
	"github.com/itispx/goimobiliar/consts"
	"github.com/itispx/goimobiliar/endereco/uf"
	"github.com/itispx/goimobiliar/erros"
	"github.com/itispx/goimobiliar/session"
)

type ActionInput struct {
	CodCondominio int `json:"CodCondominio,omitempty"` // *Código do condomínio.
}

type RunMultiInput consts.RunMultiInput[ActionInput]
type RunMultiOutput consts.RunMultiOutput[RunOutput]

func RunMulti(input *RunMultiInput) (*RunMultiOutput, error) {
	output := make(RunMultiOutput, 0, len(input.Entries))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, entry := range input.Entries {
		if input.Parallel {
			wg.Add(1)
			go func(entry *consts.RunMultiInputEntry[ActionInput]) {
				defer wg.Done()
				mu.Lock()
				output = append(output, runMultiHandler(entry))
				mu.Unlock()
			}(entry)
		} else {
			output = append(output, runMultiHandler(entry))
		}
	}

	if input.Parallel {
		wg.Wait()
	}

	return &output, nil
}

func runMultiHandler(input *consts.RunMultiInputEntry[ActionInput]) *consts.RunMultiOutputEntry[RunOutput] {
	outputEntry := consts.RunMultiOutputEntry[RunOutput]{
		ImobId: input.ImobId,
	}

	sess, err := session.NewSession(&session.NewInput{
		Endpoint: input.Endpoint,
		ImobId:   input.ImobId,
		UserId:   input.UserId,
		UserPass: input.UserPass,
	})
	if err != nil {
		msg := err.Error()

		outputEntry.Success = false
		outputEntry.Error.Message = msg

		return &outputEntry
	}

	defer sess.EndSession()

	handlerOutput, err := Run(&RunInput{
		Session:     sess,
		ActionInput: &input.Input,
	})
	if err != nil {
		msg := err.Error()

		outputEntry.Success = false
		outputEntry.Error.Message = msg

		return &outputEntry
	}

	outputEntry.Success = true
	outputEntry.Data = *handlerOutput

	return &outputEntry
}

type RunInput HandlerInput
type RunOutput RequestResponseBody

func Run(input *RunInput) (*RunOutput, error) {
	if input.Session == nil {
		return nil, erros.ErrBaseInvalida
	}

	handlerOutput, err := handler(&HandlerInput{
		Session:     input.Session,
		ActionInput: input.ActionInput,
	})
	if err != nil {
		return nil, err
	}

	return (*RunOutput)(handlerOutput.Body), nil
}

type HandlerInput struct {
	Session *session.Session
	*ActionInput
}

type HandlerOutput struct {
	*RequestResponse
}

type Request struct {
	Header *RequestHeader `json:"Header,omitempty"`
	Body   *RequestBody   `json:"Body,omitempty"`
}

type RequestHeader struct {
	SessionId string `json:"SessionId,omitempty"`
	Action    string `json:"Action,omitempty"`
}

type RequestBody struct {
	*ActionInput
}

type RequestResponse struct {
	Header *RequestResponseHeader `json:"Header,omitempty"`
	Body   *RequestResponseBody   `json:"Body,omitempty"`
}

type RequestResponseHeader struct {
	SessionId string `json:"SessionId,omitempty"`
	Action    string `json:"Action,omitempty"`
	Status    string `json:"Status,omitempty"`
	Error     bool   `json:"Error,omitempty"`
}

type RequestResponseBody struct {
	CodCondominio        int                         `json:"CodCondominio,omitempty"`        // Código do condomínio.
	NomeCondominio       string                      `json:"NomeCondominio,omitempty"`       // Nome do condomínio.
	CNPJ                 int                         `json:"CNPJ,omitempty"`                 // CNPJ do condomínio.
	TotalFracao          float64                     `json:"TotalFracao,omitempty"`          // Total das frações das economias.
	TotaldeBlocos        int                         `json:"TotaldeBlocos,omitempty"`        // Total de blocos do condomínio.
	DiaVencimentoDoc     int                         `json:"DiaVencimentoDoc,omitempty"`     // Dia de vencimento do boleto de condomínio.
	UltimaCompetenciaDoc string                      `json:"UltimaCompetenciaDoc,omitempty"` // Competência do último boleto gerado no formato YYYYMM.
	CodBlocoBase         string                      `json:"CodBlocoBase,omitempty"`         // Bloco base/principal do condomínio.
	Ativo                string                      `json:"Ativo,omitempty"`                // Indica se está ativo.
	DataInicioAdm        string                      `json:"DataInicioAdm,omitempty"`        // Data do início da administracao.
	EnderecoPrincipal    string                      `json:"EnderecoPrincipal,omitempty"`    // Endereço principal do condomínio.
	Cidade               string                      `json:"Cidade,omitempty"`               // Cidade do endereço.
	UF                   uf.UF                       `json:"UF,omitempty"`                   // Sigla da Unidade Federativa do endereço.
	Assessor             string                      `json:"Assessor,omitempty"`             // Identificação do usuário.
	AssessorNome         string                      `json:"AssessorNome,omitempty"`         // Nome do assessor/gestor.
	LojaNome             string                      `json:"LojaNome,omitempty"`             // Nome da loja/agência.
	BloqueioPagamento    string                      `json:"BloqueioPagamento,omitempty"`    // Marcação de bloqueio de pagamento.
	DataDistrato         string                      `json:"DataDistrato,omitempty"`         // Data de encerramento.
	Categoria            categoria.Categoria         `json:"Categoria,omitempty"`            // Tipo do condominio.
	Classificacao        classificacao.Classificacao `json:"Classificacao,omitempty"`        // Classificação do condominio (aba 'contrato' da tela de cadastro).
	Blocos               []*bloco.Bloco              `json:"Blocos,omitempty"`               // Array no qual cada elemento é um objeto ("Bloco" no XML) que possui os campos do objeto "bloco"
	CodAdvogadoInad      int                         `json:"CodAdvogadoInad,omitempty"`      //	Código do Advogado Inadimplente.
	NomeAdvogadoInad     string                      `json:"NomeAdvogadoInad,omitempty"`     //	Nome do Advogado Inadimplente.
	HonorarioDias        int                         `json:"HonorarioDias,omitempty"`        //	Número de dias a partir do vencimento do boleto para incidência de honorários.
	HonorarioPercentual  float64                     `json:"HonorarioPercentual,omitempty"`  //	Percentual de honorários a ser aplicado sobre o total do boleto.
}

func handler(input *HandlerInput) (*HandlerOutput, error) {
	request := Request{
		Header: &RequestHeader{
			SessionId: input.Session.SessionId,
			Action:    "CONDOM_CONDOMINIO_CONSULTAR",
		},
		Body: &RequestBody{
			input.ActionInput,
		},
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", input.Session.Endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = erros.CheckResponseError(&byteBody)
	if err != nil {
		return nil, err
	}

	var requestResponse RequestResponse
	if err := json.Unmarshal(byteBody, &requestResponse); err != nil {
		return nil, err
	}

	handlerOutput := HandlerOutput{
		RequestResponse: &requestResponse,
	}

	return &handlerOutput, nil
}
