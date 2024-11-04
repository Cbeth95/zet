# Zettelkasten
The idea is you take notes on a 'slip' and put them
in a crate. You would then link a slip to another by giving it a number (that's at least
the original idea.) This core number(?) would be it's root and then every
zettel relating to that core would have that core number.(This one idea of linking.
There are lots of other) These slips would be "atomic"
meaning one topic per slip.

You can then organize these slips by grouping them in any way and then making them into
permanent notes which are also slips. It's kind of hard to imagine since the texts on this
stuff is purposefully vague so you can create your own system.

The idea is to kind of create a 'second brain' or network of information that you can
reference back to. It is a personal tool for learning, so it doesn't adhere to any
institutional rules. Luhmann (the creator of this method) wrote 40 books in his lifetime.
40 Books.
Forty.
Books.
And apparently hundreds of academic papers. That level of productivity is chased by many by
way of this method. Books. Four tens of them. Ten books four times. I don't think I've even beat forty
games in my life.

The basic idea is you have a collection of notes all containing exactly one thought or idea (an Atomic note).
Each note has a unique.
If you are referencing a another note, you link that unique ID.

Luhmann had system of doing this:

    1. A topic
        1a.  A thought about a topic
            1a1. A thought about that thought
        1b.  Another thought about topic 1
    2. Another topic

With technology, you would use hyerlinks of those UIDs in order to link each thought, or
you can use tags to organize those thoughts into certain topics.

Either way, you are creating a web of personal knowledge that you can go back to and review.
These notes are atomic in order to break complex problems into simpler ideas that can be connected together.
It will take practice, but I think this way can help information retention for that very fact.

# Prototype

The fantasy is cool and all but I need to break it down into simpler steps:

1. Have a main command 'zet'
    
    This main command would open a neovim buffer with the title of [timestamp]
    Write the thought down, save and then exit.
    
2. The program needs a way of __linking__
    If I just wanted the current thought to link to previous thought, maybe just
    add a -p tag to the main command?

    But if I wanted to link a zettel that I don't know the unique ID of, what then?
    Maybe just tags would work, but I want this to be automatically organized.
    
    1.  Create a UID directory through arguments(zet 1123) or command tags(-n[ew]) and then
        each subsequent zet will be in that topic until a new topic is made.
    2.  Create a named directory that's then used as a zet argument (zet elephants)
        Then each zet will still have a UID, but it will be automatically grouped with the
        argument that was passed.
3. The zet needs to be __reviewed__
    As talked about previously, we need a way to review groups of zettels
    
    1. zet -r\[eview\] elephants

    Could print a buffer where a combination of zets is compiled. Which kind of defeates the purpose

    2. zet -b\[browse\] elephants

    Could open a Telescope window that has a list of that topic/tag. Then we can link
    accordingly inside each zet.
