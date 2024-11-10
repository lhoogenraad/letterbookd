"use client"

import OpenLibrarySearchForm from "./openLibrarySearch";
import api from "../../util/api/api";
import notify from 'util/notify/notify';
import { useState } from "react";

export default function OpenLibrary(){
	const [loading, setLoading] = useState(false);
	const [searchSuccess, setSearchSuccess] = useState(false);

	const submit = async (values:any) => {
		setLoading(true);
		try {
			const res = await api.openlibrary.searchForBookOL(values);
			console.log(res.data);
			setSearchSuccess(true);
		} catch (err) {
			 notify.error({message: err.response?.data?.Message ?? err.message})
		} finally {
			setLoading(false);
		}
	}

	if (searchSuccess) {
		return <div> great success </div>
	} else {
		return <OpenLibrarySearchForm submitSearch={submit} loading={loading} />
	}
}
