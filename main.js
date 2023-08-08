const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject).
then(result=>{

    goWasm.run(result.instance);

    document.getElementById("btn").addEventListener("click", (e)=>{
        document.getElementById("span").innerHTML += getHtml(" yanny")
    })
})