self.addEventListener("install", (e) =>
    e.waitUntil(
        caches
            .open("n.doolan.dev-pwa")
            .then((c) =>
                c.addAll([
                    "/",
                    "/resume",
                    "/projects?host=github&host=bgg&host=cults3d",
                    "/manifest.json",
                    "/favicon.ico",
                    "/content/resume.md",
                    "/styles/styles.css",
                    "/images/default.webp",
                    "/images/profile-192.webp",
                    "/images/profile-512.webp",
                    "/images/profile-792.webp",
                    "/images/icons/favicon-32x32.png",
                    "/images/icons/favicon-16x16.png",
                    "/images/icons/pwa-192x192.png",
                    "/images/icons/pwa-512x512.png",
                ]),
            ),
    ),
);
self.addEventListener("fetch", (e) =>
    e.respondWith(caches.match(e.request).then((r) => r || fetch(e.request))),
);
