import React from "react";
import Anagram from "./Anagram";
import { nanoid } from "nanoid";

export default function TopAnagramList(props) {
    let list = <div>No Anagrams yet!</div>
    if (props.list) {
        list = props.list.map(item => (
            <Anagram count={item.count} word={item.word} anagram={item.anagram} key={"anagram-" + nanoid()} />
        ));
    }

    return (
        <ul className="divide-y divide-gray-300">
            {list}
        </ul>
    );
}
