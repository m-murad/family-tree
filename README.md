# family-tree

## running the tool
 - clone this repository on your machine `git clone https://github.com/m-murad/family-tree`
 - cd to the root directory of the project `cd family-tree`
 - install the cli tool `go install`
 - run the tool `family-tree --help`
 
## commands
 - To add a member to the famliy
     `family-tree member <YOUR_NAME>`
     
       family-tree member AleX

 - To add a relation
      `family-tree relation <SOME_RELATION>`
      
       family-tree relation Mother

 - To define a relation between 2 people
      `family-tree relate <FIRST_NAME> <RELATION> <SECOND_NAME>`
      
       family-tree relate Alex Mother MARTHA

 - To get count of relation of a person
      `family-tree count <NAME> <RELATION>`
      
       family-tree count ALEX MOTHER

 - To get list of relations
      `family-tree list <NAME> <RELATION>`
      
       family-tree list ALEX MOTHER
       
 - To clear all members and relations
      
       family-tree list ALEX MOTHER
