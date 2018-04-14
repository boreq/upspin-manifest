# upspin-manifest
Upspin manifest helps you create manifests listing all files accessible to the
specified users. That way your friends will easily know which files are
accessible to them and you only have to provide them with a single link
to the manifest file.

## Requirements
`upspin` executable present in the `$PATH`.

## Example manifest file
Prepare the configuration file and optional header files in upspin:

    $ upspin get me@example.com/.manifest.yaml
    path: me@example.com
    manifests:
        me@example.com/share/someone/manifest:
            header: me@example.com/.manifest_headers/someone
            users: 
                - someone@example.com
        me@example.com/share/someone/manifest_noheader:
            users: 
                - someone@example.com

    $ upspin get me@example.com/.manifest_headers/someone
    Hello! This is a list of files that you can access.

## Example usage
Run `upspin-manifest` to create the specified files:

    $ upspin-manifest run me@example.com/.manifest.yaml

In the provided example the following files would be created:

    me@example.com/share/someone/manifest
    me@example.com/share/someone/manifest_noheader
