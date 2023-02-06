package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Mostrar el menú
		fmt.Println("\nSelecciona una opción:")
		fmt.Println("1. Convertir texto a Markdown")
		fmt.Println("2. Convertir texto a Jekyll")
		fmt.Println("3. Salir")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			convertToMarkdown()
		case 2:
			convertToJekyll()
		case 3:
			fmt.Println("Adiós")
			return
		default:
			fmt.Println("Opción inválida, por favor intente de nuevo")
		}
	}
}

func convertToMarkdown() {
	// Leer la entrada de texto desde la consola
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	// Convertir la entrada a Markdown
	markdown := toMarkdown(input)

	// Imprimir el resultado
	fmt.Println("\nResultado en Markdown:\n")
	fmt.Println(markdown)
}

func toMarkdown(input string) string {
	// Aplicar las reglas de formato de Markdown a la entrada de texto

	// Encabezados
	input = strings.Replace(input, "\n# ", "\n#", -1)
	for i := 6; i >= 1; i-- {
		level := strings.Repeat("#", i)
		input = strings.Replace(input, "\n"+level+" ", "\n"+level+" ", -1)
	}

	// Negritas
	input = strings.Replace(input, "*", "**", -1)

	// Cursivas
	input = strings.Replace(input, "_", "__", -1)

	// Listas no numeradas
	input = strings.Replace(input, "\n- ", "\n- ", -1)

	// Listas numeradas
	input = strings.Replace(input, "\n1. ", "\n1. ", -1)

	// Enlaces
	input = strings.Replace(input, "[", "[", -1)
	input = strings.Replace(input, "](", "](", -1)

	return input
}

func convertToJekyll() {
	// Leer la entrada de texto desde la consola
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	// Convertir la entrada a Jekyll
	jekyll := toJekyll(input)

	// Imprimir el resultado
	fmt.Println("\nResultado en Jekyll:\n")
	fmt.Println(jekyll)
}

func toJekyll(input string) string {
	// Aplicar las etiquetas necesarias para los encabezados, listas, negritas, cursivas y enlaces según la sintaxis de Jekyll

	// Encabezados
	for i := 6; i >= 1; i-- {
		level := strings.Repeat("#", i)
		input = strings.Replace(input, "\n"+level+" ", "\n<h"+strconv.Itoa(i)+">", -1)
		input = strings.Replace(input, "\n"+level, "</h"+strconv.Itoa(i)+">\n", -1)
	}

	// Negritas
	input = strings.Replace(input, "**", "<strong>", -1)
	input = strings.Replace(input, "**", "</strong>", -1)

	// Cursivas
	input = strings.Replace(input, "__", "<em>", -1)
	input = strings.Replace(input, "__", "</em>", -1)

	// Listas no numeradas
	input = strings.Replace(input, "\n- ", "\n<ul>\n<li>", -1)
	input = strings.Replace(input, "\n-", "</li>\n</ul>\n", -1)

	// Listas numeradas
	input = strings.Replace(input, "\n1. ", "\n<ol>\n<li>", -1)
	input = strings.Replace(input, "\n1.", "</li>\n</ol>\n", -1)

	// Enlaces
	input = strings.Replace(input, "[", "<a href=", -1)
	input = strings.Replace(input, "](", ">", -1)
	input = strings.Replace(input, ")", "</a>", -1)

	return input
}

