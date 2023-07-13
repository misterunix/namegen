# namegen

Thank you for looking at this little utility I call namegen.

namegen can generate male/female, middle initial, and last name/surname.

Each option has a different dataset depending on which option is selected.

The current options as of this writing are:  
**-h** help  
**-newdb** generates a new database  
**-sl** the number of surnames to load from the surname frequency set. Only used when creating a new database.  
**-c num** how many names to generate.  
**-p** user the frequency tables  
**-f** generate female first names  
**-m** generate male first names  
**-mi** pick a middle initial randomly  
**-l** pick a last name  


----
Older readme

The start of a name generator. 

The code is very messy but working so far:

Generate first name (female and male) based on just a random selection or based on the percentage frequency of the name.

Last name both random and frequency based work.

When generating a new DB the 152000+ surnames in the frequency list can be limited to any vale from all to, whatever is needed.   
The generic list of lastnames as 15000+ names.  

All text files are located in storage.
The database _sqlite_ is in the direcory db/ from the root of where the program is run from.

A count of generated names can be set as well.

Use the -h to see the current list of options.

Added README.md to all directories so they will store in git.