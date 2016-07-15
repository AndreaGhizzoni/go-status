# default option to make 
default :
	@if [ -d "deploy" ]; then rm -rf deploy; fi
	@mkdir deploy
	@cd deploy && go build ..
	@cp -r template deploy/
