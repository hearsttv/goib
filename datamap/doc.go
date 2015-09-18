package datamap

/*
	This package holds the models that map the Lakana structure to
	MediaOS structure. Its main structure is listed under Receiver,
	that is used as a union struct of all struct types to enable
	an easy JSON conversion.

	NOTES: (read this before change anything!)
	All shared common properties are abstracted in two internal types:
	- bareBase
	- fullBase
	and needs to be unmarshal by its respective methods. Note that the json tags
	are used also as a documentation reference than just a decoder hint,
	since the properties inside the specialized struct may have its name
	changed by the context. (standard set by someone else and I followed)
	So, when navigating around, look at the json field to know exactly
	what that property means.

	If a struct uses some substruct inside it (namespaced as the main
	struct), it can be found in the same file of the main struct.
	e.g.: ImageURL can be found inside image.go file
	      VideoFlavor can be found inside video.go file

	There are some properties in structs placed there just for
	backward compatibility's sake, but they're all "marked" with
	a comment. The idea is to share the same model struct between
	the migration and the goib api. So it's easy to  just change
	the endpoint later. Is that turn out to be true, I'll remove
	the backward compatibility thing, don't worry.

	There are some structs inside the receiver.go file that the
	goib API uses but are not present in the migration data
	structure. So, they're not separated in specific files as
	they are seen as subtypes for now. The only exception for
	that is Closings (that had a more complex layout and got
	separated in favor of a better organization)

*/
