<div>
<h1>Factory Method:</h1>

<h2>Definition:</h2>
<p>The Factory Method is a creational design pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created. It is used to instantiate objects without exposing the creation logic to the client and refer to the newly created object through a common interface.</p>

<h2>Implementation:</h2>
<p>In the Factory Method pattern, we create an interface or abstract class for creating an object but let the subclasses decide which class to instantiate. The Factory Method lets a class defer instantiation to subclasses. This approach promotes loose coupling and enhances code maintainability and scalability.</p>

<h2>How it works:</h2>
<p>The Factory Method works by defining a method in the base class for creating an object, which the derived classes override to specify the type of object that will be created. This method ensures that the client code uses the common interface for object creation, without needing to know the concrete class being instantiated.</p>

<h2>What can we solve with the factory method:</h2>
<h3>A real world problem:</h3>
<p>Imagine a logistics company that needs to handle different types of transport methods like trucks, ships, and planes. Each transport method has its specific implementation and characteristics. The logistics company needs a way to create instances of these transport methods without tightly coupling the client code to the specific transport classes.</p>

<h3>Solution implemented with the factory method:</h3>
<p>The Factory Method can be used to solve this problem by creating a TransportFactory interface with a method `createTransport()`. Subclasses such as TruckFactory, ShipFactory, and PlaneFactory would implement this interface to create instances of Truck, Ship, and Plane, respectively. The client code would use the TransportFactory interface to create the required transport method without knowing the specific class behind it, thus promoting flexibility and scalability.</p>

<h2>Some pros and cons</h2>
<h3>Pros:</h3>
<ul>
<li>Promotes code reusability and scalability by adhering to the Open/Closed Principle.</li>
<li>Encapsulates object creation, reducing the dependency on specific implementations.</li>
<li>Enhances code maintainability by centralizing object creation logic.</li>
</ul>

<h3>Cons:</h3>
<ul>
<li>Can increase the complexity of the code with the addition of multiple factory classes.</li>
<li>May lead to a proliferation of classes, which can be harder to manage and understand.</li>
<li>Sometimes overkill for simple object creation tasks where a simple constructor might suffice.</li>
</ul>

<h2>When to use it and when not:</h2>
<h3>When to use it:</h3>
<ul>
<li>When the exact type and dependencies of the objects are not known until runtime.</li>
<li>When you want to provide a library or framework where the users can extend and provide their implementations.</li>
<li>When you need to centralize and encapsulate the object creation logic to make the code more maintainable.</li>
</ul>

<h3>When not to use it:</h3>
<ul>
<li>When the object creation is straightforward and does not involve any complex logic or dependencies.</li>
<li>When there is no significant variation in the objects being created.</li>
<li>When the overhead of creating and maintaining multiple factory classes outweighs the benefits of flexibility and decoupling.</li>
</ul>

<h4>Grafic representation: </h4>
<img src="./assets/Screenshot 2024-07-04 at 10.02.15.png"/>
</div>
