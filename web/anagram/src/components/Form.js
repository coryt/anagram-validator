import React, { useState } from "react";

export default function Form(props) {
    const [showResult, setShowResult] = React.useState(false)
    const [firstWord, setFirstWord] = useState('');
    const [secondWord, setSecondWord] = useState('');

    function handleFirstWordChange(e) {
        setFirstWord(e.target.value);
    }

    function handleSecondWordChange(e) {
        setSecondWord(e.target.value);
    }

    const submit = e => {
        e.preventDefault()
        fetch('http://0.0.0.0:4000/api/anagrams/check', {
            method: 'POST',
            body: JSON.stringify({ firstWord, secondWord }),
            headers: { 'Content-Type': 'application/json' },
        })
            .then(res => res.json())
            .then(json => {
                let result = ""
                if (json.valid === true) {
                    result = <div className="py-3 px-5 mb-4 bg-green-100 text-green-900 text-sm rounded-md border border-green-200" role="alert"><strong>Nice</strong> those words are anagrams!</div>
                } else {
                    result = <div className="py-3 px-5 mb-4 bg-red-100 text-red-900 text-sm rounded-md border border-red-200" role="alert"><strong>Sorry :(</strong> those words are not anagrams!</div>
                }
                setShowResult(result)
                setFirstWord("");
                setSecondWord("");
                props.onSubmit();
            })

    }

    return (
        <form className="m-4" onSubmit={submit}>
            <div className="flex my-2">
                <div className="flex mx-4">
                    <div className="relative rounded-md shadow-sm">
                        <input type="text" id="first-word-input" name="text" className="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker" placeholder="First Word" autoComplete="off" value={firstWord} onChange={handleFirstWordChange} />
                    </div>
                </div>
                <div className="flex mx-4">
                    <div className="relative rounded-md shadow-sm">
                        <input type="text" id="second-word-input" name="text" className="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker" placeholder="Second Word" autoComplete="off" value={secondWord} onChange={handleSecondWordChange} />
                    </div>
                </div>
            </div>
            <div className="mt-8 flex justify-center">
                <div className="inline-flex rounded-md bg-white shadow">
                    <button type="submit" className="text-gray-700 font-bold py-2 px-6">Validate</button>
                </div>
            </div>
            <div className="mt-8 flex justify-center">
                {showResult}
            </div>
        </form>
    );
}