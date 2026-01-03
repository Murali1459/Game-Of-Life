# Game of Life — Terminal Experiment

This is an experimental implementation of **Conway’s Game of Life**, built to run entirely in the terminal.

It’s not meant to be a perfect or optimized version.  
It’s meant to be something you can **watch, poke, and disturb**.

This project is part of my experimental series **Undefined Behavior** — a space for curiosity-driven builds, strange ideas, and systems explored without a fixed outcome.

---

## What’s different here?

### Cells have age
Cells aren’t just alive or dead.

Each living cell carries an **age**, which increases every generation it survives.  
Age isn’t shown numerically—instead, it leaks into the display:

. → ░ → ▒ → █


New cells barely register.  
Older cells grow denser, heavier, harder to ignore.

Time leaves a visible trace.

---

### It runs in the terminal
No GUI. No canvas. No abstractions hiding the rules.

The terminal is treated as the medium, not a limitation.

---

### You can interfere
This isn’t a closed system.

Click anywhere in the grid and a new live cell (age = 1) is injected into the simulation.  
Not as control—but as **perturbation**.

You’re not directing the system.  
You’re nudging it and watching how it responds.

---

## Rules (mostly Conway’s Game of Life)

- Any live cell with fewer than 2 or more than 3 neighbors dies
- Any dead cell with exactly 3 neighbors becomes alive
- Live cells age every generation they survive
- Dead cells have age 0

Internally, age can grow without bound.  
Visually, it saturates.

---

## Controls

- **Mouse click**: spawn a live cell at the cursor
- **Any key**: quit

---

## Why this exists

This project isn’t about novelty or usefulness.

It’s about:
- Exploring simple systems with time as a first-class concept
- Seeing how stability reacts to disturbance
- Rediscovering the joy of building things just to understand them

No AI assistance was used while building this.

---

## Notes

- Window resizing resets the simulation
- The system is intentionally imperfect
- Rough edges are part of the experiment

---
