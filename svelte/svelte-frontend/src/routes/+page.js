
export async function load({data}){
    var result = await fetch("http://localhost:81")
    var bs = result.text();
    return{
        test:bs,
    }
}
