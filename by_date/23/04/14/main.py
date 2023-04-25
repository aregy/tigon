WS = { "\t": True, " " : True, "\n" : True }

def reverseWords(s : str) -> str:
    t = list(s)
    d = dict()
    if t[0]
    return ''.join(t)

def test_reverseWords():
    phrase = "words are here"
    assert reverseWords(phrase) == "here are words"

if __name__ == "__main__":
    while True:
        s : str = input(": ")
        print(reverseWords(s)l)
