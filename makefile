# default option to make 
default :
	@if [ -d "_obj" ]; then rm -rf _obj; fi
	@mkdir _obj
	@cd _obj && go build ..
	@if [ -d "deploy" ]; then rm -rf deply; fi
	@mkdir deploy
	@mv _obj/go-status deploy/
	@cp -r template deploy/
