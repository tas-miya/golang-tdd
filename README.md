# golang-tdd

### 📓 **Lesson 01 - Notes**

### 🟨 **Go Basics**

- **`package main`**: Entry point for executables.
- **`func main()`**: Program's starting function.
- **`import "fmt"`**: Imports standard output formatting tools.
- **`fmt.Println()`**: Prints output to the console.
- **Functions**: Defined using `func` keyword. Can return values like `string`.

---

### ✅ **Testing in Go**

- Go uses built-in testing tools—no setup required.
- Tests live in `_test.go` files.
- Test functions must:

  - Start with `Test`
  - Accept one param: `t *testing.T`

- Use `t.Errorf()` for failing test output.
- Use `%q` to print quoted strings in error messages.

---

### 🧪 **Test-Driven Development (TDD) Steps**

1. Write a failing test first.
2. Write minimal code to pass the test.
3. Refactor with tests as safety net.
4. Repeat.

---

### 🧱 **Improving Test Quality**

- Use **subtests** with `t.Run("desc", func(t *testing.T) {...})`.
- Extract repeated assertions into helper:

  ```go
  func assertCorrectMessage(t testing.TB, got, want string)
  ```

- Mark helpers with `t.Helper()` to clean up error tracebacks.

---

### 📦 **Go Modules**

- Required from Go 1.16+.
- Initialize with:
  `go mod init example.com/hello`
- Creates a `go.mod` file for dependency management.

---

### ⚙️ **Refactoring**

- Use **constants** to remove "magic strings":

  ```go
  const englishHelloPrefix = "Hello, "
  ```

- Improve readability and reuse.

---

### 🌍 **Adding Language Support**

- Modify `Hello(name, language string)` to return localized greeting.
- Use `switch` for clarity over multiple `if`s:

  ```go
  switch language {
  case "Spanish":
      return "Hola, " + name
  }
  ```

- Handle empty `name` with default "World".

---

### 📉 **Extract Helper Logic**

- Move logic into private function:

  ```go
  func greetingPrefix(language string) string
  ```

- Use named return values for clarity.

---

### 🌐 **Documentation Tools**

- `go doc fmt`: View local docs.
- `pkgsite`: Browse docs locally via web UI.

---

### 🔁 **Development Discipline**

- Always see your test **fail first**.
- Stick to the TDD cycle for clarity and robustness.
- Refactor tests as well as production code.

---

### 🏁 **Concepts Covered**

- Go syntax: functions, `if`, `const`, `switch`
- Variables: short declaration `:=`
- Tests: built-in, lightweight, standard
- TDD: Fail → Pass → Refactor loop

---

### 📓**Lesson 02 - Notes**

### 📁 **Project Structure**

- Each package should have its own directory
- go.mod needs to be in the root directory

---

### 🧪 **Test-Driven Development Using Convey**

- Used Convey instead of standard testing framework
- Usual TDD steps: test → fail → write minimal code until the test passes → refactor

---

### 💭 **Use Comments for Documentation**

```go
  // Add takes two integers and return their sum
  func Add(a, b int) int {
  	return a + b
  }
```

---

### 📘 **Testable Examples**

- They are used to write examples of the usage of the function
- They are written in the \_test.go file

```go
  func ExampleAdd() {
  Convey("Example of Add Function", func() {
    got := Add(1, 2)
    So(got, ShouldEqual, 3)
    })
  }
```
