# import tkinter as tk
from tkinter import * # type: ignore
from prefixsuggest import Trie
from tkinter import Event
from typing import Tuple # type: ignore

data_list = [
    "apple", "apricot", "avocado", "banana", "blackberry", "blueberry", "boysenberry", "breadfruit", "cantaloupe", "cherry",
    "clementine", "coconut", "cranberry", "currant", "date", "dragonfruit", "durian", "elderberry", "fig", "gooseberry",
    "grape", "grapefruit", "guava", "honeydew", "huckleberry", "jackfruit", "jujube", "kiwi", "kumquat", "lemon",
    "lime", "longan", "loquat", "lychee", "mandarin", "mango", "mangosteen", "marionberry", "melon", "mulberry",
    "nectarine", "orange", "papaya", "passionfruit", "peach", "pear", "persimmon", "pineapple", "plantain", "plum",
    "pomegranate", "pomelo", "quince", "raspberry", "rambutan", "redcurrant", "starfruit", "strawberry", "tangerine", "tamarind",
    "tamarillo", "ugli fruit", "watermelon", "yuzu", "bilberry", "blackcurrant", "blood orange", "canary melon", "cherimoya", "cloudberry",
    "damson", "feijoa", "finger lime", "goji berry", "honeyberry", "jabuticaba", "kiwano", "langsat", "mamey sapote", "miracle fruit",
    "nance", "olive", "pawpaw", "physalis", "pineberry", "prickly pear", "pulasan", "quandong", "rose apple", "salak",
    "santol", "sapodilla", "soursop", "sugar apple", "surinam cherry", "tamarillo", "tayberry", "white currant", "white sapote", "wolfberry"
]



root = Tk()
root.title('Trie Data Strutucre use for debouce! ')
root.geometry("500x300")


# update list box
def update_box(data: list[str]):
    # clear the list box
    listBox.delete(0, END)

    for d in data:
        listBox.insert(END, d)


# add on click
def add_on_click(event: Event):
    selection_indices: Tuple[int, ...] = listBox.curselection()  # type: ignore
    if not selection_indices:
        return
    
    index: int = selection_indices[0] # type: ignore
    selection: str = listBox.get(index) # type: ignore
    
    searchBox.delete(0, "end")
    searchBox.insert(0, str(selection)) # type: ignore

# checking typed word
def check_typed_word(event: Event):
    # get whatever is typed
    typed = searchBox.get()

    if typed == "":
        data = data_list
    data = trie.search(typed)

    update_box(data)




Label(root, text="Start searching...", font=("Helvetica", 18)).pack(pady=10)

searchBox = Entry(root, font=("Helvetica", 16))
searchBox.pack(padx=20, pady=10)

listBox = Listbox(root, font=("Helvetica", 18))
listBox.pack(padx=10, pady=20)

trie = Trie()

for dl in data_list:
    trie.insert(dl)

update_box(data=data_list)

# bind to the listbox
listBox.bind("<<ListboxSelect>>", add_on_click)

# search box binding
searchBox.bind("<KeyRelease>", check_typed_word)

root.mainloop()