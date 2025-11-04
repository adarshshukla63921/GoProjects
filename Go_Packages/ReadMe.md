**To create your own custom pacakge in Go**

1) Create a subfolder, inside your module, each of these folders represent one package.
2) Inside the folder, create your own **.go** file.
3) Ensure to not use **main** as your package name, use something different (ideally use your folder name itself to avoid confusion).
4) **Exported functions or variable** must have names that begin with capital letter, otherwise you won't be able to access them outside the package itself

In Go, each file should follow a certain package.
