import React from 'react'

export class NewChannelModal extends React.Component {
    constructor(props){
        super(props)
    }

    render(){
        return (
            <div className="modal fade" id="NewChannelModal" tabIndex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
                <div className="modal-dialog" role="document">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="exampleModalLongTitle">Modal title</h5>
                            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                    <div className="modal-body">
                        <input type="text" className="form-control" id="NewChannelModal" placeholder="Enter channel name" onChange={this.props.onChannelIDChange}/>
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary" data-dismiss="modal">Close</button>
                        <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={this.props.onCreateChannel}>Save changes</button>
                    </div>
                    </div>
                </div>
            </div>
        )
    }
}