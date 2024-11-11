"use client"

import OpenLibrarySearchForm from "./openLibrarySearch";
import OpenLibraryBook from "./openLibraryBook";
import api from "../../util/api/api";
import notify from 'util/notify/notify';
import { useState } from "react";

export default function OpenLibrary({close}){
	const [loading, setLoading] = useState(false);
	const [searchSuccess, setSearchSuccess] = useState(false);
	const [olBook, setOLBook] = useState({});

	const submitSearch = async (values:any) => {
		setLoading(true);
		try {
			const res = await api.openlibrary.searchForBookOL(values);
			console.log(res.data);
			setOLBook(res.data)
			setSearchSuccess(true);
		} catch (err) {
			 notify.error({message: err.response?.data?.Message ?? err.message})
		} finally {
			setLoading(false);
		}
	}

	const submitAdd = async (id: string) => {
		setLoading(true);
		try {
			const res = await api.openlibrary.attemptToAddBookFromOL({ id });
			notify.info({message: "Successfully added " + res.data.Title});
			close();
		} catch (err) {
			 notify.error({message: err.response?.data?.Message ?? err.message})
		} finally {
			setLoading(false);
		}
	}

	if (searchSuccess) {
		return <OpenLibraryBook submitAdd={submitAdd} loading={loading} olBook={olBook} />
	} else {
		return <OpenLibrarySearchForm submitSearch={submitSearch} loading={loading} />
	}
}
