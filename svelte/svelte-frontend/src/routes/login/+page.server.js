import { redirect } from '@sveltejs/kit';



export function load({ locals }) {
	if (locals.user) throw redirect(302, '/notes');
}

/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({cookies,request}) => {
      const data = await request.formData();
      const email = data.get('email');
      const password = data.get('password');
      const response = await fetch('http://localhost:81/login', {
        method: 'POST',
        body: JSON.stringify({"Email": email, "Password": password}),
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        }

      });if(!response){
        return{
            success: false,
        }
      }
      let dataPost;
      try{
       dataPost = await response.json();
      }catch(error){
        console.log(error);
        return{
            success: false,
        }
      }
    //   let parsed = JSON.parse(total);
    //   console.log(dataPost['token']);

      

        cookies.set('jwt', dataPost,{path: '/',
    maxAge: dataPost['expiration']} )
    //   return{
    //     success: true,
    //     data: dataPost,
        
    //   }
    throw redirect(301, '/notes');
      
    
      
}
    
  };