install-npm:
	@echo "--> Installing Node dependencies"
	cd ui && npm install

test-js:
	@echo "--> Building static assets"
	cd ui && node_modules/.bin/webpack -p
	@echo "--> Running JavaScript tests"
	cd ui && npm run test
	@echo ""

lint-js:
	@echo "--> Linting javascript"
	cd ui && node_modules/.bin/eslint  --config .eslintrc --ext .jsx,.js {tests/js,app}
	@echo ""
