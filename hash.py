import hashlib


def main():
    password = input("Escribe la contraseña que quieres cifrar : ")
    m = hashlib.sha256()
    m.update(password.encode("utf-8"))

    return m.hexdigest()


if __name__ == "__main__":
    hash = main()
    print(hash)
