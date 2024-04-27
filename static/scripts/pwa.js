document.addEventListener("readystatechange", () => {
    const isLightMode = window.matchMedia(
        "(prefers-color-scheme: light)",
    ).matches;
    document.body.classList.toggle("dark", !isLightMode);
});
window.addEventListener("load", async () => {
    if ("serviceWorker" in navigator) {
        try {
            await navigator.serviceWorker.register("scripts/service_worker.js");
        } catch (e) {
            console.log("Service Worker registration failed");
        }
    }
});
