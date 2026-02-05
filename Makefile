.PHONY: help tag push-tag push-tags release

help:
	@echo Alvos disponíveis:
	@echo   tag        - Criar uma tag git (uso: make tag VERSION=v1.0.0)
	@echo   push-tag   - Enviar uma tag específica para o remoto (uso: make push-tag VERSION=v1.0.0)
	@echo   push-tags  - Enviar todas as tags para o remoto
	@echo   release    - Criar e enviar uma tag (uso: make release VERSION=v1.0.0)

tag:
	git tag -a $(VERSION) -m "Lançamento $(VERSION)"
	@echo Tag $(VERSION) criada!

push-tag:
	git push origin $(VERSION)
	@echo Tag $(VERSION) enviada!

push-tags:
	git push --tags
	@echo Todas as tags enviadas!
	
release:
	git tag -a $(VERSION) -m "Lançamento $(VERSION)"
	git push origin $(VERSION)
	@echo Lançamento $(VERSION) completo!