package ctapag_condominio_notafiscal_importar

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/itispx/goimobiliar/consts"
	"github.com/itispx/goimobiliar/erros"
	"github.com/itispx/goimobiliar/session"
)

var ACTION = "CTAPAG_CONDOMINIO_NOTAFISCAL_IMPORTAR"

type ActionInput struct {
	CodCondominio        int     `json:"CodCondominio,omitempty"`        // *Código do condomínio do lançamento (se origem for 'C').
	CodBloco             string  `json:"CodBloco,omitempty"`             // Código do bloco do lançamento (se origem for 'C').
	CodFornecedor        int     `json:"CodFornecedor,omitempty"`        // *Código do fornecedor do lançamento.
	DataEmissao          string  `json:"DataEmissao,omitempty"`          // *Data de emissão do lançamento (se TipoDocumento for 'N').
	DataVencimento       string  `json:"DataVencimento,omitempty"`       // *Data de vencimento do lançamento.
	TipoDocumento        string  `json:"TipoDocumento,omitempty"`        // *Tipo de documento do lançamento.
	FormaPagamento       string  `json:"FormaPagamento,omitempty"`       // Forma de pagamento do lançamento.
	CodTaxa              int     `json:"CodTaxa,omitempty"`              // *Código da taxa que classifica este lançamento.
	NumeroParcela        int     `json:"NumeroParcela,omitempty"`        // Número da parcela do lançamento. Valor default é '1'.
	TotalParcelas        int     `json:"TotalParcelas,omitempty"`        // Quantidade total de parcelas. Valor default é '1'.
	Complemento          string  `json:"Complemento,omitempty"`          // Complemento descritivo do lançamento.
	NumeroDocumento      string  `json:"NumeroDocumento,omitempty"`      // *Número do documento do fornecedor.
	ValorBruto           float64 `json:"ValorBruto,omitempty"`           // *Valor bruto do documento/parcela.
	ValorServicos        float64 `json:"ValorServicos,omitempty"`        // Valor dos serviços. Se não informado, a base de cálculo será ValorBruto.
	ValorBaseCalculoIss  float64 `json:"ValorBaseCalculoIss,omitempty"`  // Base de cálculo do ISS. Se não informado, a base de cálculo será ValorServicos.
	ValorRetencaoInss    float64 `json:"ValorRetencaoInss,omitempty"`    // Valor do INSS a ser retido.
	ValorRetencaoIss     float64 `json:"ValorRetencaoIss,omitempty"`     // Valor do ISS a ser retido.
	ValorRetencaoIrf     float64 `json:"ValorRetencaoIrf,omitempty"`     // Valor do IRF a ser retido.
	ValorRetencaoFederal float64 `json:"ValorRetencaoFederal,omitempty"` // Valor da retenção federal a ser retida.
	Comissao             float64 `json:"Comissao,omitempty"`             // Valor de comissão.
	CodigoBarras         string  `json:"CodigoBarras,omitempty"`         // Código de barras do documento (* obrigatório se origem for 'B')
	PixQrCode            string  `json:"PixQrCode,omitempty"`            // QR Code.
	PrevisaoReal         string  `json:"PrevisaoReal,omitempty"`         // *Indicação de lançamento previsto ou real.
	UrlImagem            string  `json:"UrlImagem,omitempty"`            // URL para efetuar download da imagem, por exemplo "http://imagens.com.br/lancto123.pdf".
}

type RunMultiInput consts.RunMultiInput[*ActionInput]
type RunMultiOutput consts.RunMultiOutput[*RunOutput]

func RunMulti(input *RunMultiInput) (*RunMultiOutput, error) {
	output := make(RunMultiOutput, 0, len(input.Entries))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, entry := range input.Entries {
		if input.Parallel {
			wg.Add(1)
			go func(entry *consts.RunMultiInputEntry[*ActionInput]) {
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

func runMultiHandler(input *consts.RunMultiInputEntry[*ActionInput]) *consts.RunMultiOutputEntry[*RunOutput] {
	outputEntry := consts.RunMultiOutputEntry[*RunOutput]{
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
		ActionInput: input.Input,
	})
	if err != nil {
		msg := err.Error()

		outputEntry.Success = false
		outputEntry.Error.Message = msg

		return &outputEntry
	}

	outputEntry.Success = true
	outputEntry.Data = handlerOutput

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
	NumeroLancto int `json:"NumeroLancto,omitempty"` // Número do lançamento.
}

func handler(input *HandlerInput) (*HandlerOutput, error) {
	request := Request{
		Header: &RequestHeader{
			SessionId: input.Session.SessionId,
			Action:    ACTION,
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
