import React from 'react'

type Props = {
  params: {id: string}
}

const UpdateItemPage = (props: Props) => {
  const {id} = props.params
  return (
    <>
      <div>UpdateItemPage</div>
      <div>{id}</div>
    </>
  )
}

export default UpdateItemPage