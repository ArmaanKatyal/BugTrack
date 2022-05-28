import React from 'react'
import { useParams } from 'react-router-dom'

function ProjectTickets(props) {
    const params = useParams();
  return (
    <div>{params.id}</div>
  )
}

export default ProjectTickets