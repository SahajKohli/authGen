# authGen
A GO package that will allow the user to generate a API and other Auth 

I wanted to create a package that would allow my GoLang and NodeJS projects generate API keys so I created this package. There is another package that wrote in JS that has the logic so that I can do the same thing for my backend NodeJS projects.

### API Key Generation:

Its relatively simple function that will return a API in this format: xxxxxx-xxxxxx-xxxx-xxxxxx. Uses the current rand.Intn to do most of the heavy lifting here. The ceiling is set at 127, which is the amount of ASCII chars, so I basically split into half from 0 - 63 and 64 - 127. Where anything from 0 - 63 is divded by 9 and then converted into a int. Anything higher then 63 (not including 63) is then divided by 2 and converted into a lowercase alphabetic char. 

### JWT Generation:

TODO

Contribution: If you want to make some changes, feel free to do a PR! Always open to advice.

Random Comment: I didn't expect to see such a stark difference between JS and Golang in terms of how much effort it would take to rewrite this in GO. After trying to do it in a similar fashion I realized that the fact that JS lets you take a callback function makes things much easier and cleaner. 
