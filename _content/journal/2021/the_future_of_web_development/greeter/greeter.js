"use strict";
exports.__esModule = true;
var Greeter = /** @class */ (function () {
    function Greeter(id) {
        this.el = document.getElementById(id);
        if (!this.el) {
            console.error("element id not found:", id);
            return;
        }
    }
    Greeter.prototype.Greet = function (person) {
        this.el.innerHTML = "Hello, " + person.firstName + " " + person.lastName;
    };
    return Greeter;
}());
