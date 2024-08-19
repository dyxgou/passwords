def preguntar_numero(mensaje: str):
    entrada = input(mensaje)

    if not entrada.isnumeric():
        print("Esa entrada es invalida.")
        return preguntar_numero(mensaje)

    return int(entrada)


def main():
    password = input("Escribe el texto que quieres encriptar : ")
    pasos = preguntar_numero("Escribe la cantidad de pasos : ")

    hash = ""

    for letra in password:
        uni = ord(letra) + pasos
        hash += chr(uni)

    return hash


if __name__ == "__main__":
    hash = main()
    print(hash)
