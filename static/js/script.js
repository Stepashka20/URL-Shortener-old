let url = ""
let shortUrl = ""
async function copyToClipboard(text){
    if (navigator.clipboard) {
        await navigator.clipboard.writeText(text)
    } else {
        const textArea = document.createElement('textarea')
        textArea.value = text
        document.body.appendChild(textArea)
        textArea.select()
        document.execCommand('copy')
        textArea.remove()
    }
}
document.getElementById("makeshorterBtn").addEventListener("click",async  function() {
    if (shortUrl){
        //copy to clipboard
        copyToClipboard(shortUrl)

        document.querySelector("#makeshorterBtn span").innerHTML = "Сократить";
        document.querySelector("#url").value = "";
        shortUrl = "";
        return;
    }
    //regex for url validation
    if (!/.+\..+/.exec(document.querySelector("#url").value)){
        SnackBar({
            message: "Введите корректный url",
            status: "error",
            timeout: 1000,
            position:"tr"
    
        });
        return;
    }

    
    document.getElementById("makeshorterBtn").classList.add("disabled");
    document.querySelector("#makeshorterBtn span").style.display = "none";
    document.querySelector("#makeshorterBtn .loading-spinner").style.display = "block";
    //remove last / from location.href
    try {
        const raw = await fetch(`${location.href.replace(/\/$/, "")}/getShortUrl`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                url: document.querySelector("#url").value
            })
        });
        const data = await raw.json();
        if (raw.ok){
            shortUrl = `${location.href.replace(/\/$/, "")}/${data.shortUrl}`
            document.getElementById("makeshorterBtn").classList.remove("disabled");
            document.querySelector("#makeshorterBtn span").style.display = "block";
            document.querySelector("#makeshorterBtn .loading-spinner").style.display = "none";

            document.querySelector("#makeshorterBtn span").innerHTML = "Скопировать";
            document.querySelector("#url").value = shortUrl;
            document.querySelector("#url").select();
        } else {
            SnackBar({
                message: data.error,
                status: "error",
                timeout: 1000,
                position:"tr"

            });
        }
    } catch (error) {
        SnackBar({
            message: "Ошибка сервера",
            status: "error",
            timeout: 1000,
            position:"tr"
        });
    }


})
