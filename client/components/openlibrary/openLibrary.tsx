"use client"

import OpenLibrarySearchForm from "./openLibrarySearch";
import api from "../../util/api/api";

export default function OpenLibrary(){
	const submit = async (values:any) => {
		const res = await api.openlibrary.searchForBookOL(values);
		console.log(res.data);
	}

	return (
		<OpenLibrarySearchForm submitSearch={submit} loading={false} />
	)
}
