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
}

function editor_show_js() {
    document.getElementById("editor-js").classList.toggle("show");

    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-css").classList.remove("show");
    document.getElementById("editor-scss").classList.remove("show");
}

function editor_show_css() {
    document.getElementById("editor-css").classList.toggle("show");

    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-js").classList.remove("show");
    document.getElementById("editor-scss").classList.remove("show");
}

function editor_show_scss() {
    document.getElementById("editor-scss").classList.toggle("show");

    document.getElementById("editor-html").classList.remove("show");
    document.getElementById("editor-js").classList.remove("show");
    document.getElementById("editor-css").classList.remove("show");
}

// ACE Editor Collection

var ace_editor_html;

ace_editor_html = ace.edit("ace_editor_html");
ace_editor_html.setTheme("ace/theme/monokai");
ace_editor_html.session.setMode("ace/mode/html");

var ace_editor_js;

ace_editor_js = ace.edit("ace_editor_js");
ace_editor_js.setTheme("ace/theme/monokai");
ace_editor_js.session.setMode("ace/mode/javascript");

var ace_editor_css;

ace_editor_css = ace.edit("ace_editor_css");
ace_editor_css.setTheme("ace/theme/monokai");
ace_editor_css.session.setMode("ace/mode/css");

var ace_editor_scss;

ace_editor_scss = ace.edit("ace_editor_scss");
ace_editor_scss.setTheme("ace/theme/monokai");
ace_editor_scss.session.setMode("ace/mode/scss");

// Code from E.L.P. Prasandika
// ---------------------------

// Nothing