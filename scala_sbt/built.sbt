lazy val root = (project in file(".")).settings(
  name := "hello",
  version := "1.0",
  scalaVersion := "2.13.8",
  libraryDependencies += "org.scalatest" %% "scalatest-funsuite" % "3.2.11" % Test
)
