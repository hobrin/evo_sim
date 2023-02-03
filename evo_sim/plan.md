# Evo Sim Plan.

## Priorities
* Computational speed -> low quality physics (-> low tickrate), short lifetimes, GPU?
* Complexity -> interactiveness, complex behaviours, tradeoffs -> (strategy?)
* Biodiversity -> the possibilities are N^2 where N is the amount of features, promotes local minima

The amount of entities could always be maximum supported by the (GPU?) at a time. This ensures maximum ensures maximum efficiency.

## Reproduction method. 
* **Asexual selection** reproduces an organism one for one. This is good for speciation (I think) but is bad for genetic in a population. Every organism would split up and grow out to entire new species in theory. Organisms don't like this and this is extremely slow, meaning a lot of computational complexity for nothing. This is basically a no go.
* **Sexual Selection** blends the 2 organisms into a new ofspring. The advantage is that a gene can actually spread through a population.
* **Death** might actually be mean that it is teleported and mutatated with DNA from prosperous members of the population.

## Environment
* **Day-night time cycle** would insentivize different hunting strategies and sleeping. And sleeping incentivizes building a safe home. 
* **Slowness but increased eyesight** fase would be a fase in the game where the organisms can see more in order to explore, but they could also opt in to sleep through. 
* **Hyperspeed but reduced eyesight** fase


## Organism features and properties
* **Energy** would be the ultimate currency of the organisms.
* **Speed** would be a tradeoff between energy and speed. In hunting situations it would be prioritised.
* **Leap** would affect timer, distance and energy usage.
* **Sleep** would be a mode in which an entity sleeps. The benefits of sleep could be auto balanced at runtime based on a population analysis, or it could be that an organism will get so much perforamce reduction because of insomnia that it has no other choice but to sleep.
* **Eyesight** would vary in degrees.
* **Eyedistance** would vary in distance. Tradeoff would be that it could report more false positives.
* **Block building** would be the ability to build blobs for a certain amount of time. Cost energy.
* **Projectiles** would be the ability to throw projectiles. 
* **Harnass** would be decreased mobility.
* **Regen**
* **Max health**



## Brain
* We could use a connection based brain, where the organisms would own let's say 10 connections, and then in their genes they specify with whom and which neurons they will connect the connections. 