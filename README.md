## Zettelkasten
The idea is you take notes on a 'slip' and put them
in a crate. You would then link a slip to another by giving it a number (that's at least
the original idea.) This core number(?) would be it's root and then every
zettel relating to that core would have that core number.(This one idea of linking.
There are lots of other) These slips would be "atomic"
meaning one topic per slip.

You can then organize these slips by grouping them in any way and put them in a kasten then making them into
permanent notes which are also slips. It's kind of hard to imagine since the texts on this
stuff is purposefully vague, so you can create your own system.

The basic idea is you have a collection of notes all containing exactly one thought or idea (an Atomic note).
Each note has a unique.
If you are referencing another note, you link that unique ID.

Luhmann had system of doing this:

    1. A topic
        1a.  A thought about a topic
            1a1. A thought about that thought
        1b.  Another thought about topic 1
    2. Another topic

>This is just my understanding of the system. This README will update when I understand it a little more

By making each slip exactly one thought in your own words, you retain more information
about the topic you are 'zetteling'(?) about. Notes are no longer a long page but a collection
of individual thoughts. 

## Usage
This repo is an attempt an emulation of this technique using neovim. The command `gf`
actually searches for a file under that current position of the cursor and opens it in a new buffer.
Linking each thought would be as simple as:

```markdown
A thought (Link to another thought)[ UID.md ]
```

If you're using neovim, you wouldn't be dealing with a bunch of windows/tabs but different buffers.

Eventually of course we would need to add functionality to more editors but for simplicity, I'm just using
my personal editor.

## Goals

### Have a main command 'zet'
    
    This main command would open a neovim buffer with the title of [timestamp]
    Write the thought down, save and then exit.
    
### The program needs a way of __linking__

As already discussed in the usage section, we can link a file by using a file link. If we're in the
same directory, we can just use `gf` but we can add more linking capabilities as we go.
### The zet needs to be __reviewed__

There can be a few ways to do this:

1. zet -r\[eview\] elephants

Could print a buffer where a combination of zets is compiled. 

2. zet -b\[browse\] elephants

Could open a telescope.nvim window that shows list of that topic/tag. Then we can link
accordingly inside each zet.

