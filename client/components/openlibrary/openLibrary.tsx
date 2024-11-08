"use client"

import OpenLibrarySearchForm from "./openLibrarySearch";
import api from "../../util/api/api";

export default function OpenLibrary(){
	const submit = async (values) => {
		await api.openlibrary.searchForBookOL(values);
	}

	return (
		<OpenLibrarySearchForm submitSearch={submit} loading={false} />
	)
}
