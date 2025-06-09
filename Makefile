build-page:
	 hugo -d dist --minify -s my-site --baseURL "https://konzepte-moderner-softwareentwicklung.github.io/Docu/"

run:
	 hugo server --minify -s my-site
