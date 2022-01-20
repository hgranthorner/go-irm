# go-irm

An interactive rm, for careful deleters.

## Design

```markdown
[ ] thing.txt
[X] other_thing.txt
 >  closed_directory
 v  open_directory
  [ ] File 1
  [X] File 2
```

## Spec

- [X] View files in directory
- [ ] Expand subdirectories
- [ ] Mark files for deletion
- [ ] Mark directories for deletion
- [ ] Delete marked objects
- [ ] Dry run to file
- [ ] Allow user to input regex and mark files for them