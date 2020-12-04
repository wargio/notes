package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsc92352cbca454b798c4a03bc11520803252e0d8e = "{{define \"base\"}}\n<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <link rel=\"stylesheet\" href=\"//cdn.rawgit.com/picturepan2/spectre/v0.2.14/docs/dist/spectre-icons.min.css\">\n    <link rel=\"stylesheet\" href=\"//cdn.rawgit.com/picturepan2/spectre/v0.2.14/docs/dist/spectre.min.css\">\n    {{ template \"stylesheets\" . }}\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    {{ template \"css\" . }}\n    <title>Notes</title>\n    <style type=\"text/css\">\n      .spacing-top { margin: 3px 0px 0px 0px; }\n      .spacing-left { margin: 0px 3px 0px 0px; }\n      .float-right { float: right; }\n      .display-block { display: block; }\n    </style>\n  </head>\n<body>\n  <section class=\"container grid-960 mt-10\">\n    <header class=\"navbar\">\n      <section class=\"navbar-section\">\n        <a href=\"{{.Root}}\" class=\"navbar-brand mr-10\">Notes</a>\n      </section>\n      <section class=\"navbar-section\">\n        <a href=\"{{.Root}}new\" class=\"btn btn-primary\">New Note</a>\n      </section>\n    </header>\n    {{template \"content\" .}}\n  </section>\n</body>\n{{ template \"scripts\" . }}\n</html>\n{{end}}\n{{ define \"css\" }}{{ end }}\n{{ define \"scripts\" }}{{ end }}\n{{ define \"stylesheets\" }}{{ end }}\n"
var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "{{define \"content\"}}\n<section class=\"container\">\n  <div class=\"columns\">\n    <div class=\"column\">\n      {{ $root := .Root}}\n      {{ range $Note := .NoteList }}\n        <div class=\"input-group spacing-top\">\n          <span class=\"input-group-addon spacing-left\">\n            {{ $Note.Title }}\n          </span>\n          <a class=\"btn btn-action spacing-left\" href=\"{{$root}}view/{{$Note.ID}}\">\n            <i class=\"icon icon-forward\">View</i>\n          </a>\n        </div>\n      {{end}}\n    </div>\n  </div>\n</section>\n{{end}}\n"
var _Assetsfaf8c71f8f367d08490b3fa4da14715443eb2b6f = "{{define \"content\"}}\n<section class=\"container\">\n  <div class=\"columns\">\n    <div class=\"column\">\n      <div class=\"panel\">\n        <div class=\"panel-header\">\n          <div class=\"panel-title\">{{.Title}}</div>\n        </div>\n        <div class=\"panel-body\">{{.HTML}}</div>\n        <div class=\"panel-footer\">\n          <div class=\"btn-group display-block\">\n            <a class=\"btn btn-primary mr-5\" href=\"{{.Root}}edit/{{.ID}}\">\n              <i class=\"icon icon-edit\"></i>\n              Edit\n            </a>\n            <a class=\"btn float-right\" href=\"{{.Root}}delete/{{.ID}}\">\n              <i class=\"icon icon-delete\"></i>\n              Delete\n            </a>\n          </div>\n        </div>\n      </div>\n    </div>\n  </div>\n</section>\n{{end}}\n"
var _Assets0b1feb78498d6b8e6774a1306860e6920bfe3253 = "{{define \"stylesheets\"}}\n<link rel=\"stylesheet\" href=\"//cdn.rawgit.com/sparksuite/simplemde-markdown-editor/1.11.2/dist/simplemde.min.css\">\n{{end}}\n{{define \"content\"}}\n<section class=\"container\">\n  <div class=\"columns\">\n    <div class=\"column\">\n      <form action=\"{{.Root}}save{{if .ID}}/{{.ID}}{{end}}\" method=\"POST\">\n        <div class=\"form-group\">\n          <input type=\"hidden\" name=\"id\" value=\"{{.ID}}\" />\n          <input type=\"text\" class=\"form-input\" name=\"title\" value=\"{{.Title}}\" placeholder=\"Title\"/>\n          <textarea class=\"form-input\" id=\"input-body\" name=\"body\" placeholder=\"Content\" cols=\"80\" rows=\"24\">{{printf \"%s\" .Body}}</textarea>\n        </div>\n        <div class=\"form-group\">\n          <button class=\"btn btn-primary\" type=\"submit\">Save</button>\n          <a class=\"btn btn-link\" href=\"{{.Root}}\">Cancel</a>\n        </div>\n      </form>\n    </div>\n  </div>\n</section>\n{{end}}\n{{define \"scripts\"}}\n<script src=\"//cdn.rawgit.com/jquery/jquery/3.2.1/dist/jquery.slim.min.js\"></script>\n<script src=\"//cdn.rawgit.com/sparksuite/simplemde-markdown-editor/1.11.2/dist/simplemde.min.js\"></script>\n<script>\n$(document).ready(function() {\n  var simplemde = new SimpleMDE({\n      element: $(\"#input-body\")[0],\n      autofocus: true,\n      hideIcons: [\"side-by-side\", \"fullscreen\", \"guide\"],\n      autosave: {\n          enabled: true,\n          uniqueId: \"notes\", // TODO: Make this configurable?\n          delay: 1000,\n      },\n      forceSync: true,\n      indentWithTabs: false,\n      promptURLs: true,\n      tabSize: 4,\n  });\n});\n</script>\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html", "view.html", "edit.html", "base.html"}}, map[string]*assets.File{
	"/base.html": &assets.File{
		Path:     "/base.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1594820943, 1594820943381848044),
		Data:     []byte(_Assetsc92352cbca454b798c4a03bc11520803252e0d8e),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1594818054, 1594818054967591562),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1594821003, 1594821003331429285),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}, "/view.html": &assets.File{
		Path:     "/view.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1594820946, 1594820946785157694),
		Data:     []byte(_Assetsfaf8c71f8f367d08490b3fa4da14715443eb2b6f),
	}, "/edit.html": &assets.File{
		Path:     "/edit.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1594818054, 1594818054967591562),
		Data:     []byte(_Assets0b1feb78498d6b8e6774a1306860e6920bfe3253),
	}}, "")
