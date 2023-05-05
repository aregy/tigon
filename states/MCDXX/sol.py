#!/usr/bin/env python3

# We can break up the string greedily. First we'll break up s into an array words. Then, we can use a buffer as a current string and tentatively add words to it, checking that the newly-added-to line can fit within k. If we overflow with the new word, then we flush out the current string into an array all and restart it with the new word.
#
# Notice that if any word is longer than k, then there's no way to break up the text, so we should return None. It's helpful to define a helper function that returns the length of a list of words with spaces added in between.
#
# Finally, we return all, which should contain the texts we want.
#
def break(s, k):
    words = s.split()

    if not words:
        return []

    current = []
    all = []

    for i, word in enumerate(words):
        if length(current + [word]) <= k:
            current.append(word)
        elif length([word]) > k:
            return None
        else:
            all.append(current)
            current = [word]
    all.append(current)

    return all

def length(words):
    if not words:
        return 0
    return sum(len(word) for word in words) + (len(words) - 1)
