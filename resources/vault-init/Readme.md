### Scone notes
1. Image
   1. For Injection of file image is used in session description. It can have injection_file or volume
      1. volume - system region that can use fspf, and be shared by exporting 
      2. injection_files - These are updated by CAS session if they have $$SCONE$$ in them, if the content is not provided
        contents of the file at the specified path are stored in-memory which are serverd by scone