use tauri::Builder;
use warpnet_frontend_dist::FrontendAssets;
use std::borrow::Cow;
use http::{Response, HeaderValue};
use mime_guess;

pub fn run() {
    Builder::default()
        .register_uri_scheme_protocol("app", |_ctx, request: tauri::http::Request<Vec<u8>>| {
            println!("uri path: '{}'", request.uri().path());
            let mut uri_path = request.uri().path().trim_start_matches('/');

            if uri_path.is_empty() {
                uri_path = "index.html";
            }

            // Путь к ассету: если прямой match, то берём его, иначе fallback
            let (asset_path, asset) = match FrontendAssets::get(uri_path) {
                Some(asset) => (uri_path, Some(asset)),
                None => ("index.html", FrontendAssets::get("index.html")),
            };

            match asset {
                Some(asset) => {
                    let mime = mime_guess::from_path(asset_path).first_or_octet_stream();
                    let body = match asset.data {
                        Cow::Borrowed(b) => b.to_vec(),
                        Cow::Owned(b) => b,
                    };

                    Response::builder()
                        .header("Content-Type", HeaderValue::from_str(mime.as_ref()).unwrap())
                        .body(body)
                        .unwrap()
                }
                None => Response::builder()
                    .status(404)
                    .body(Vec::new())
                    .unwrap(),
            }
        })
        .run(tauri::generate_context!())
        .expect("failed to run tauri app");
}
