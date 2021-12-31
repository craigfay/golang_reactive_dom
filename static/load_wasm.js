
// Loading and executing webassembly modules

const go = new Go();
let mod, inst;


WebAssembly.instantiateStreaming(
  fetch("dist/main.wasm", { cache: 'no-cache' }),
  go.importObject,
)
  .then(async (result) => {
    mod = result.module;
    inst = result.instance;
    await go.run(inst);
  });
