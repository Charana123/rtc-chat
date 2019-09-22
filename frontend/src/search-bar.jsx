import React from 'react'

export default function SearchBar(props) {
    return (
        <div className="has-search">
            <span className="fa fa-search form-control-feedback"></span>
            <input 
                type="text" 
                className="form-control"
                placeholder="Search"
                onChange={props.onChange}
                onKeyPress={props.onKey}
            />
        </div>
    )
}