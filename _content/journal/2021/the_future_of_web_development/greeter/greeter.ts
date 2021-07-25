import { Person, Student } from "../models/models"

class Greeter {
	private el!: HTMLElement | null

	constructor(id: string) {
		this.el = document.getElementById(id)
		if (!this.el) {
			console.error("element id not found:", id)
			return
		}
	}

	Greet(person: Person) {
		this.el.innerHTML = "Hello, " + person.firstName + " " + person.lastName
	}
}
