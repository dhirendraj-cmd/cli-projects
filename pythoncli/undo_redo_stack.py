def main():
    undo: list[str] = []
    redo: list[str] = []

    HELP = "Commands: write <text> | undo | redo | clear | show | help"
    

    while True:
        user_input = input("> ").strip()
        user_input = user_input.lower()

        if user_input.startswith("write "):
            add = user_input[len("write "):]
            if add:
                undo.append(add)
                redo.clear()
            continue

        elif user_input == "undo".lower():
            if undo:
                redo.append(undo.pop())
            else:
                print("Nothing to undo!".upper())
            continue

        elif user_input == "redo".lower():
            if redo:
                undo.append(redo.pop())
            else:
                print("Nothing to redo!!".upper())
            continue

        elif user_input == "clear".lower():
            undo.clear()
            redo.clear()

        elif user_input == "show".lower():
            current = " ".join(undo)
            print(f"Current text: '{current}' ")

        elif user_input == "help".lower():
            print("Below are the commands needs to be used!!")
            print(HELP)

        elif user_input == "exit".lower():
            break

        else:
            print("Unknown Command, Please type 'help' to show the list of commands")



if __name__=="__main__":
    main()
