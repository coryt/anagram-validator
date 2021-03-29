import React from "react";

export default function Anagram(props) {
    return (
        <li className="p-4">
            <div className="flex flex-row justify-center">
                <div className="mr-2">
                    ({props.count})
                </div>
                <div className="mr-2">
                    {props.word}
                </div>
                <div>
                    {props.anagram}
                </div>
            </div>
        </li>
    );
}
