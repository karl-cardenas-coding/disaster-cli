license:
	@echo "Applying license headers..."
	 copywrite headers


opensource:
	@echo "Checking for open source licenses"
	~/go/bin/go-licenses report github.com/karl-cardenas-coding/disaster-cli --template=documentation/open-source.tpl > documentation/open-source.md 