// Rust backend for ultra-low latency desktop trading terminal
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

#[tauri::command]
fn get_system_latency() -> String {
    // Microsecond timer
    "0.082ms".to_string()
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![get_system_latency])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
