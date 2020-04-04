require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

let processCodeBtn = document.querySelector("#process-code");
processCodeBtn.onclick = () => {
    // send javascript code over to Go interpreter somehow
    let codeArea = $('textarea#code-area').val();
    console.log(codeArea);
}
$(() => {
    console.log("webpage has loaded")
});
