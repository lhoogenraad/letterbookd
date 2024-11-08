"use client"

import OpenLibrarySearchForm from "./openLibrarySearch";
import api from "../../util/api/api";
import { useState } from "react";

export default function OpenLibrary(){
	const [loading, setLoading] = useState(false);

	const submit = async (values:any) => {
		setLoading(true);
		try {
			const res = await api.openlibrary.searchForBookOL(values);
			console.log(res.data);
		} finally {
			setLoading(false);
		}
	}

	return (
		<OpenLibrarySearchForm submitSearch={submit} loading={loading} />
	)
}
