// Java Script for Codinoc Theme
//
// Written by:
// - N.L.M. Nishshanke
// - E.C.N. Nandasiri
// - E.L.P. Prasandika

// Code from N.L.M. Nishshanke
// ---------------------------

// Nothing

// Code from E.C.N. Nandasiri
// --------------------------

// Show menu when click menu item button

function logged_nav_menu() {
    document.getElementById("dropdown-menu").classList.toggle("show");

    document.getElementById("dropdown-settings").classList.remove("show");
    document.getElementById("dropdown-notification").classList.remove("show");
}

function logged_nav_settings() {
    document.getElementById("dropdown-settings").classList.toggle("show");

    document.getElementById("dropdown-menu").classList.remove("show");
    document.getElementById("dropdown-notification").classList.remove("show");
}

function logged_nav_notification() {
    document.getElementById("dropdown-notification").classList.toggle("show");

    document.getElementById("dropdown-menu").classList.remove("show");
    document.getElementById("dropdown-settings").classList.remove("show");
}

// Change editor windows

function editor_show_html() {
    document.getElementById("editor-html").classList.toggle("show");
    document.getElementById("editor-js").classList.remove("show");
    document.getElementById("editor-css").classList.remove("show");
    document.getElementById("editor-scss").classList.remove("show");

    document.getElementById("editor-btn-html").classList.toggle("activated");
    document.getElementById("editor-btn-js").classList.remove("activated");
    document.getElementById("editor-btn-css").classList.remove("activated");
    document.getElementById("editor-btn-scss").classList.remove("activated");
}

function editor_show_js() {
    document.getElementById("editor-js").classList.toggle("show");
    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-css").classList.remove("show");
    document.getElementById("editor-scss").classList.remove("show");

    document.getElementById("editor-btn-js").classList.toggle("activated");
    document.getElementById("editor-btn-html").classList.remove("activated");
    document.getElementById("editor-btn-css").classList.remove("activated");
    document.getElementById("editor-btn-scss").classList.remove("activated");
}

function editor_show_css() {
    document.getElementById("editor-css").classList.toggle("show");
    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-js").classList.remove("show");
    document.getElementById("editor-scss").classList.remove("show");

    document.getElementById("editor-btn-css").classList.toggle("activated");
    document.getElementById("editor-btn-html").classList.remove("activated");
    document.getElementById("editor-btn-js").classList.remove("activated");
    document.getElementById("editor-btn-scss").classList.remove("activated");
}

function editor_show_scss() {
    document.getElementById("editor-scss").classList.toggle("show");
    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-js").classList.remove("show");
    document.getElementById("editor-css").classList.remove("show");

    document.getElementById("editor-btn-scss").classList.toggle("activated");
    document.getElementById("editor-btn-html").classList.remove("activated");
    document.getElementById("editor-btn-js").classList.remove("activated");
    document.getElementById("editor-btn-css").classList.remove("activated");
}

// Code from E.L.P. Prasandika
// ---------------------------

// Nothing
