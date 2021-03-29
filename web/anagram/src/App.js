import React, { useEffect, useState } from "react";
import Form from "./components/Form";
import TopAnagramList from "./components/TopAnagramList";

function App(token) {
  const [topAnagramList, setTopAnagramList] = useState();

  function fetchTopAnagrams() {
    fetch('http://0.0.0.0:4000/api/anagrams/top', {
      method: 'GET',
    })
      .then(res => res.json())
      .then(json => {
        updateTopAnagrams(json)
      })
  }

  function updateTopAnagrams(topAnagrams) {
    setTopAnagramList(topAnagrams);
  }

  useEffect(() => {
    let mounted = true;
    fetch('http://0.0.0.0:4000/api/anagrams/top', {
      method: 'GET',
    })
      .then(res => res.json())
      .then(json => {
        if (mounted) {
          updateTopAnagrams(json)
        }
      })
    return () => mounted = false;
  }, [])
  return (
    <div className="anagramapp flex items-center justify-center h-screen">
      <div className="max-w-screen-lg bg-indigo-500 shadow-2xl rounded-lg mx-auto text-center py-12 mt-4">
        <h2 className="text-3xl leading-9 font-bold tracking-tight text-white sm:text-4xl sm:leading-10">
          Anagram Validator
        </h2>
        <Form onSubmit={fetchTopAnagrams} />
        <div className="mt-8">
          <h2 className="text-xl leading-9 font-bold tracking-tight text-white sm:text-xl sm:leading-10">
            Top 10 Anagrams Searched
          </h2>
          <div className="flex justify-center">
            <div className="bg-white shadow-xl rounded-lg w-1/2">
              <TopAnagramList list={topAnagramList} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
