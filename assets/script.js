document.getElementById("urlForm").addEventListener("submit", (event) => {
  event.preventDefault();

  const urlInput = document.getElementById("url");
  const url = urlInput.value;

  if (!url) {
    urlInput.focus();
    return;
  }

  fetch(`/shorten?url=${url}`)
    .then((res) => res.text())
    .then((url) => {
      const shortUrlOutput = document.getElementById("shortUrlOutput");
      const shortUrlOutputLabel = document.getElementById(
        "shortUrlOutputLabel"
      );
      shortUrlOutput.value = url;
      shortUrlOutput.style.display = "block";
      shortUrlOutputLabel.style.display = "block";
    });
});
