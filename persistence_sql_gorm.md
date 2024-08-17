# gorm [[{persistence.sql]]
Problems with Gorm:
https://www.reddit.com/r/golang/comments/umkgk3/gorm_is_a_bad_idea/

* Exiting the Vietnam (of Gorm):
  https://alanilling.com/exiting-the-vietnam-of-programming-our-journey-in-dropping-the-orm-in-golang-3ce7dff24a0f

  ...  Alternatives:
  We bucketed projects in the Go community along two lines:
      Code-generating SQL at runtime (example: squirrel)
      Generating application code at compile time(examples: jet, sqlc)

   We looked closely at two code generators: jet and sqlc, ultimately
  selecting sqlc. With jet you write SQL within your application as a
  DSL. But because it generates code it goes a step beyond what a
  runtime SQL generator like squirrel offers. Models and fields and are
  first-class referenceable types, rather than requiring string
  interpolation, which avoids the need to grep through code in an audit
  process when you want to make changes.
[[}]]

