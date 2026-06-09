# cli util for file creating from template

how using

```bash
./gotpl -s source -tf template -fs '' -es ''
```

-fs = '⟦⟦' - символ открытия шаблонной строки
-es = '⟧⟧' - символ закрытия шаблонной строки

путь до файла в котором меняем шаблонные строки source.js
```js
(function(e){(function(){var t={}.hasOwnProperty;function n(){for(var a=⟦⟦key1⟧⟧,i=0;i<arguments.length;i++)}})})
```


template
```config.yml
key1: value1
key2: value2
```


Компилить под linux 

```bash
GOOS=linux GOARCH=amd64 go build -o gotpl ./cmd/gotpl
```